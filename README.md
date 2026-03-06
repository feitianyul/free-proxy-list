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

最后更新：2026-03-06 07:52:33 UTC（2026-03-06 15:52:33 UTC+8）

**代理总数：61**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 125.128.12.144:3128 | ✓ 1811ms | ✓ 1923ms | ✓ 775ms | ✓ 1249ms | ✓ 904ms | http |
| 205.209.118.30:3138 | ✓ 1361ms | ✓ 1617ms | 否 | 否 | ✓ 922ms | http |
| 61.72.221.234:3128 | ✓ 1400ms | 否 | ✓ 1435ms | 否 | ✓ 1995ms | http |
| 35.225.22.61:80 | ✓ 298ms | ✓ 1062ms | 否 | 否 | ✓ 1003ms | http |
| 130.36.36.29:443 | ✓ 978ms | 否 | ✓ 1086ms | ✓ 1006ms | ✓ 947ms | http |
| 8.219.97.248:80 | ✓ 1661ms | 否 | ✓ 1343ms | 否 | ✓ 1743ms | http |
| 168.235.110.63:3128 | ✓ 1072ms | 否 | 否 | ✓ 1337ms | ✓ 1036ms | http |
| 46.249.103.192:443 | ✓ 884ms | 否 | ✓ 1746ms | ✓ 1897ms | 否 | http |
| 165.225.242.47:12963 | ✓ 583ms | ✓ 866ms | ✓ 967ms | ✓ 1045ms | ✓ 948ms | http |
| 165.225.50.18:10086 | ✓ 672ms | ✓ 976ms | ✓ 726ms | ✓ 1344ms | ✓ 1250ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1530ms | ✓ 1430ms | ✓ 920ms | ✓ 666ms | http |
| 121.128.121.54:3128 | ✓ 1722ms | 否 | ✓ 1787ms | 否 | ✓ 853ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1240ms | ✓ 1960ms | ✓ 1557ms | http |
| 125.128.12.14:3128 | ✓ 1677ms | 否 | 否 | ✓ 1246ms | ✓ 1789ms | http |
| 107.174.80.186:3128 | ✓ 1605ms | ✓ 1380ms | ✓ 879ms | 否 | 否 | http |
| 61.72.110.54:3128 | ✓ 1675ms | 否 | 否 | ✓ 1834ms | ✓ 1592ms | http |
| 14.56.107.244:3128 | ✓ 1821ms | ✓ 1686ms | ✓ 1054ms | 否 | ✓ 1040ms | http |
| 91.107.175.112:10801 | ✓ 870ms | 否 | 否 | ✓ 1822ms | ✓ 1788ms | http |
| 138.124.53.25:7443 | ✓ 478ms | ✓ 1970ms | 否 | 否 | ✓ 1584ms | http |
| 172.212.68.37:3128 | ✓ 229ms | 否 | ✓ 879ms | ✓ 1524ms | ✓ 1013ms | http |
| 61.72.221.194:3128 | ✓ 1512ms | 否 | ✓ 912ms | ✓ 1907ms | 否 | http |
| 61.72.221.94:3128 | 否 | 否 | ✓ 1879ms | ✓ 1892ms | ✓ 1197ms | http |
| 116.80.82.222:3172 | ✓ 1727ms | 否 | 否 | ✓ 1973ms | ✓ 1807ms | http |
| 192.166.82.55:1080 | ✓ 1272ms | 否 | ✓ 1730ms | ✓ 1191ms | ✓ 1244ms | http |
| 91.193.240.157:9877 | ✓ 798ms | 否 | ✓ 860ms | 否 | ✓ 1384ms | http |
| 61.72.110.94:3128 | 否 | 否 | ✓ 1694ms | ✓ 1879ms | ✓ 1666ms | http |
| 101.43.255.96:80 | ✓ 1419ms | ✓ 1445ms | 否 | ✓ 1765ms | ✓ 1075ms | http |
| 89.185.85.138:1080 | ✓ 575ms | ✓ 1950ms | ✓ 1209ms | ✓ 1960ms | ✓ 1149ms | http |
| 103.215.36.88:19137 | ✓ 1513ms | ✓ 1495ms | ✓ 1325ms | 否 | ✓ 1488ms | http |
| 160.250.4.245:1 | ✓ 1932ms | 否 | ✓ 1765ms | ✓ 1442ms | ✓ 1150ms | http |
| 160.250.5.22:1 | ✓ 1583ms | 否 | ✓ 1754ms | ✓ 1447ms | ✓ 1293ms | http |
| 39.104.201.40:7890 | ✓ 1775ms | ✓ 1335ms | ✓ 1688ms | ✓ 1349ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1153ms | ✓ 1383ms | ✓ 1112ms | ✓ 1365ms | ✓ 1053ms | http |
| 120.92.212.16:7890 | ✓ 1890ms | ✓ 1330ms | ✓ 1040ms | ✓ 1587ms | ✓ 1367ms | http |
| 14.56.177.44:3128 | ✓ 1705ms | ✓ 1822ms | ✓ 1434ms | ✓ 1793ms | 否 | http |
| 5.75.196.26:40000 | ✓ 516ms | 否 | ✓ 847ms | 否 | ✓ 966ms | http |
| 74.48.78.224:2080 | ✓ 1184ms | ✓ 1721ms | ✓ 1148ms | 否 | 否 | http |
| 38.180.2.107:3128 | ✓ 829ms | 否 | ✓ 1536ms | 否 | ✓ 1699ms | http |
| 103.215.36.88:16988 | ✓ 1266ms | ✓ 1381ms | 否 | ✓ 1552ms | ✓ 1195ms | http |
| 45.136.198.40:3128 | ✓ 962ms | 否 | ✓ 1557ms | 否 | ✓ 1614ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1632ms | ✓ 1625ms | ✓ 1184ms | http |
| 104.129.202.19:10227 | ✓ 508ms | ✓ 1408ms | ✓ 1288ms | ✓ 1199ms | ✓ 873ms | http |
| 98.98.26.42:9443 | ✓ 1352ms | 否 | ✓ 1208ms | 否 | ✓ 1656ms | http |
| 165.225.51.24:11647 | ✓ 611ms | ✓ 1694ms | ✓ 756ms | ✓ 936ms | ✓ 661ms | http |
| 165.225.51.24:3128 | ✓ 601ms | ✓ 1952ms | ✓ 536ms | ✓ 866ms | ✓ 710ms | http |
| 121.230.8.45:1080 | ✓ 1331ms | ✓ 1801ms | 否 | ✓ 1422ms | ✓ 1308ms | http |
| 121.230.8.49:1080 | ✓ 1329ms | ✓ 1526ms | ✓ 1315ms | ✓ 1730ms | ✓ 1364ms | http |
| 69.48.179.20:3128 | ✓ 244ms | 否 | ✓ 871ms | ✓ 1384ms | 否 | http |
| 39.98.86.246:8118 | ✓ 1032ms | 否 | ✓ 1981ms | 否 | ✓ 1868ms | http |
| 59.8.203.55:80 | ✓ 1778ms | ✓ 1583ms | ✓ 1769ms | 否 | 否 | http |
| 165.225.210.43:11602 | ✓ 1138ms | 否 | ✓ 546ms | ✓ 1055ms | ✓ 931ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1369ms | 否 | ✓ 1332ms | ✓ 1345ms | http |
| 103.215.36.88:16895 | ✓ 1351ms | ✓ 1765ms | 否 | ✓ 1436ms | ✓ 1418ms | http |
| 103.215.36.88:15556 | ✓ 1134ms | ✓ 1628ms | ✓ 1250ms | ✓ 1358ms | 否 | http |
| 1.225.116.115:1080 | 否 | ✓ 1167ms | ✓ 1100ms | ✓ 1085ms | 否 | http |
| 120.79.99.232:8099 | ✓ 1371ms | ✓ 1659ms | ✓ 1445ms | ✓ 1616ms | ✓ 1336ms | http |
| 103.215.36.88:19288 | 否 | 否 | ✓ 1835ms | ✓ 1391ms | ✓ 1784ms | http |
| 45.140.147.155:1081 | ✓ 958ms | 否 | ✓ 798ms | ✓ 1274ms | ✓ 868ms | http |
| 154.37.208.132:30000 | ✓ 956ms | ✓ 1805ms | 否 | ✓ 1723ms | ✓ 1700ms | http |
| 103.166.185.54:3128 | ✓ 1521ms | 否 | ✓ 1161ms | ✓ 1407ms | ✓ 1194ms | http |
| 188.132.141.249:443 | ✓ 1119ms | 否 | ✓ 1320ms | 否 | ✓ 1950ms | http |

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
