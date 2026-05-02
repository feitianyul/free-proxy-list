**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-05-02 15:44:18 UTC（2026-05-02 23:44:18 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 206.206.126.177:2412 | ✓ 818ms | 否 | ✓ 1806ms | ✓ 1420ms | ✓ 1100ms | http |
| 113.160.132.26:8080 | ✓ 1515ms | 否 | ✓ 1005ms | ✓ 1698ms | ✓ 1398ms | http |
| 218.108.131.186:17890 | 否 | 否 | ✓ 1209ms | ✓ 1762ms | ✓ 1315ms | http |
| 62.60.231.71:56608 | ✓ 683ms | 否 | ✓ 1200ms | 否 | ✓ 1077ms | http |
| 45.167.124.71:999 | ✓ 881ms | 否 | 否 | ✓ 1681ms | ✓ 1370ms | http |
| 107.150.41.226:18080 | ✓ 1485ms | 否 | ✓ 647ms | ✓ 1590ms | ✓ 1891ms | http |
| 20.127.128.70:8080 | ✓ 1620ms | 否 | ✓ 532ms | ✓ 1847ms | ✓ 1591ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1921ms | ✓ 1784ms | ✓ 1502ms | http |
| 217.76.245.80:999 | ✓ 1560ms | 否 | ✓ 1281ms | ✓ 1578ms | ✓ 1354ms | http |
| 45.153.231.229:8080 | ✓ 868ms | 否 | ✓ 1560ms | 否 | ✓ 1948ms | http |
| 154.12.231.32:80 | ✓ 576ms | 否 | ✓ 1700ms | 否 | ✓ 1077ms | http |
| 212.58.132.5:8888 | ✓ 1173ms | 否 | ✓ 1352ms | ✓ 1488ms | ✓ 1292ms | http |
| 72.11.150.178:6005 | ✓ 323ms | 否 | 否 | ✓ 1690ms | ✓ 954ms | http |
| 20.78.26.206:8561 | ✓ 1608ms | 否 | ✓ 1647ms | ✓ 1863ms | 否 | http |
| 20.78.118.91:8561 | ✓ 1613ms | 否 | ✓ 1669ms | ✓ 1846ms | 否 | http |
| 149.51.42.10:8080 | ✓ 749ms | ✓ 1287ms | 否 | ✓ 1296ms | 否 | http |
| 47.85.51.197:1080 | ✓ 240ms | 否 | ✓ 587ms | ✓ 1409ms | ✓ 897ms | http |
| 86.104.72.220:1081 | ✓ 906ms | ✓ 1232ms | ✓ 165ms | ✓ 1559ms | ✓ 853ms | http |
| 103.35.190.69:1082 | ✓ 901ms | ✓ 995ms | ✓ 1230ms | 否 | ✓ 826ms | http |
| 86.104.72.220:1082 | ✓ 911ms | ✓ 1234ms | ✓ 1192ms | ✓ 1136ms | 否 | http |
| 148.230.4.241:999 | ✓ 958ms | ✓ 1660ms | ✓ 644ms | 否 | 否 | http |
| 80.92.204.47:1081 | ✓ 1138ms | ✓ 1869ms | ✓ 1934ms | 否 | ✓ 1236ms | http |
| 91.184.241.12:443 | 否 | ✓ 1790ms | ✓ 1217ms | ✓ 1980ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1737ms | 否 | ✓ 1534ms | ✓ 1253ms | ✓ 949ms | http |
| 213.111.146.36:18080 | ✓ 1386ms | ✓ 1925ms | ✓ 1404ms | 否 | ✓ 1420ms | http |
| 43.133.44.89:8888 | ✓ 1430ms | 否 | ✓ 1436ms | ✓ 1377ms | ✓ 1443ms | http |
| 120.92.108.86:7890 | ✓ 1540ms | 否 | 否 | ✓ 1865ms | ✓ 1434ms | http |
| 103.193.144.43:3125 | ✓ 1408ms | 否 | 否 | ✓ 1477ms | ✓ 1450ms | http |
| 103.193.145.22:8082 | ✓ 1425ms | 否 | ✓ 1511ms | ✓ 1616ms | ✓ 1799ms | http |
| 101.32.244.83:8080 | ✓ 1101ms | 否 | ✓ 1033ms | ✓ 1356ms | ✓ 1389ms | http |
| 121.43.196.213:8222 | ✓ 991ms | ✓ 1163ms | ✓ 928ms | ✓ 1243ms | ✓ 1048ms | http |
| 121.43.196.210:8222 | ✓ 1051ms | ✓ 1233ms | ✓ 945ms | ✓ 1189ms | ✓ 992ms | http |
| 86.104.72.219:1081 | ✓ 1424ms | 否 | ✓ 198ms | 否 | ✓ 1048ms | http |
| 160.238.65.8:3128 | ✓ 1599ms | 否 | ✓ 1588ms | ✓ 1879ms | 否 | http |
| 160.238.65.3:3128 | ✓ 492ms | ✓ 1841ms | ✓ 476ms | ✓ 1420ms | ✓ 1049ms | http |
| 160.238.65.9:3128 | ✓ 453ms | 否 | ✓ 483ms | ✓ 1335ms | ✓ 1033ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 723ms | ✓ 1687ms | ✓ 1386ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 723ms | ✓ 1687ms | ✓ 1385ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 723ms | ✓ 1696ms | ✓ 1388ms | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 734ms | ✓ 1688ms | ✓ 1389ms | http |
| 160.238.65.6:3128 | ✓ 526ms | 否 | ✓ 456ms | ✓ 1354ms | ✓ 1039ms | http |
| 117.236.124.166:3128 | ✓ 1363ms | 否 | ✓ 1283ms | 否 | ✓ 1966ms | http |
| 120.92.212.16:8890 | ✓ 1452ms | ✓ 1234ms | 否 | 否 | ✓ 1086ms | http |
| 49.156.44.114:8080 | ✓ 1669ms | 否 | ✓ 1352ms | ✓ 1771ms | ✓ 1480ms | http |
| 149.51.42.10:3128 | ✓ 406ms | ✓ 1267ms | 否 | ✓ 1465ms | 否 | http |
| 152.42.177.32:8888 | ✓ 1063ms | 否 | ✓ 1176ms | ✓ 1409ms | ✓ 1351ms | http |
| 160.238.65.2:3128 | ✓ 532ms | ✓ 1921ms | 否 | 否 | ✓ 1134ms | http |
| 3.101.133.120:80 | ✓ 370ms | ✓ 1176ms | ✓ 1926ms | ✓ 1380ms | ✓ 896ms | http |
| 47.77.216.82:1080 | ✓ 546ms | 否 | ✓ 1105ms | ✓ 801ms | ✓ 601ms | http |
| 160.238.65.5:3128 | ✓ 657ms | 否 | ✓ 905ms | 否 | ✓ 1705ms | http |
| 94.131.118.39:1081 | ✓ 981ms | ✓ 1871ms | ✓ 1356ms | 否 | 否 | http |
| 130.61.174.200:1080 | 否 | 否 | ✓ 497ms | ✓ 1330ms | ✓ 1147ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 817ms | ✓ 1131ms | ✓ 787ms | http |
| 20.210.39.153:8561 | 否 | 否 | ✓ 971ms | ✓ 941ms | ✓ 739ms | http |
| 152.70.91.193:40000 | ✓ 1722ms | 否 | 否 | ✓ 1531ms | ✓ 1576ms | http |
| 94.158.219.111:3128 | ✓ 798ms | 否 | ✓ 987ms | 否 | ✓ 1625ms | http |
| 92.119.127.208:6005 | ✓ 1052ms | ✓ 1646ms | ✓ 1470ms | ✓ 1932ms | ✓ 1812ms | http |
| 160.238.65.4:3128 | ✓ 620ms | 否 | ✓ 1801ms | ✓ 1391ms | ✓ 1854ms | http |
| 86.104.72.219:1082 | ✓ 386ms | 否 | ✓ 219ms | ✓ 1710ms | 否 | http |
| 101.32.243.189:80 | ✓ 1429ms | 否 | ✓ 1348ms | ✓ 1509ms | ✓ 1422ms | http |
| 61.52.131.172:8443 | ✓ 943ms | ✓ 1257ms | ✓ 968ms | ✓ 1270ms | ✓ 1068ms | http |
| 103.172.70.173:8080 | ✓ 1166ms | 否 | ✓ 1984ms | ✓ 1654ms | ✓ 1433ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
