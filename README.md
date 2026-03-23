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

最后更新：2026-03-23 16:58:29 UTC（2026-03-24 00:58:29 UTC+8）

**代理总数：58**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 20.210.39.153:8561 | ✓ 1322ms | 否 | ✓ 640ms | ✓ 947ms | ✓ 737ms | http |
| 20.78.26.206:8561 | ✓ 1440ms | 否 | ✓ 599ms | ✓ 889ms | ✓ 719ms | http |
| 147.161.210.140:8800 | ✓ 1322ms | 否 | ✓ 815ms | ✓ 1348ms | ✓ 1147ms | http |
| 147.161.239.240:8800 | ✓ 634ms | 否 | ✓ 1348ms | ✓ 1687ms | ✓ 1616ms | http |
| 167.103.34.108:8800 | ✓ 1511ms | 否 | ✓ 1756ms | ✓ 1658ms | 否 | http |
| 45.167.124.52:8080 | ✓ 932ms | ✓ 1881ms | ✓ 1667ms | 否 | ✓ 1648ms | http |
| 113.160.132.26:8080 | ✓ 1476ms | 否 | 否 | ✓ 1955ms | ✓ 1121ms | http |
| 20.78.118.91:8561 | ✓ 579ms | ✓ 905ms | ✓ 1397ms | ✓ 1346ms | ✓ 1496ms | http |
| 186.148.180.46:999 | ✓ 962ms | 否 | ✓ 521ms | ✓ 1918ms | ✓ 1358ms | http |
| 167.103.31.122:8800 | ✓ 1397ms | 否 | ✓ 1501ms | ✓ 1779ms | ✓ 1596ms | http |
| 35.225.22.61:80 | ✓ 509ms | ✓ 1160ms | ✓ 1056ms | ✓ 1102ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1832ms | ✓ 1319ms | ✓ 1497ms | ✓ 1305ms | ✓ 1213ms | http |
| 120.92.212.16:8890 | ✓ 980ms | ✓ 1797ms | 否 | ✓ 1487ms | ✓ 1421ms | http |
| 101.43.127.100:8877 | ✓ 923ms | 否 | ✓ 1341ms | ✓ 1157ms | 否 | http |
| 116.80.65.75:3172 | ✓ 1555ms | 否 | ✓ 1596ms | ✓ 1948ms | 否 | http |
| 121.230.9.104:1080 | ✓ 1509ms | ✓ 1531ms | ✓ 1733ms | 否 | 否 | http |
| 181.78.44.63:999 | ✓ 757ms | ✓ 1648ms | ✓ 1428ms | ✓ 1780ms | ✓ 1339ms | http |
| 83.219.250.8:62920 | ✓ 1207ms | 否 | ✓ 1684ms | 否 | ✓ 1761ms | http |
| 47.101.149.27:9010 | ✓ 1895ms | 否 | ✓ 1438ms | ✓ 1872ms | 否 | http |
| 185.115.74.185:8080 | ✓ 1062ms | ✓ 1796ms | ✓ 1958ms | 否 | 否 | http |
| 38.145.208.240:8448 | 否 | ✓ 1207ms | ✓ 442ms | ✓ 815ms | ✓ 909ms | http |
| 2.83.243.148:7777 | ✓ 1257ms | ✓ 1726ms | 否 | ✓ 1849ms | ✓ 1934ms | http |
| 103.113.70.189:1081 | 否 | 否 | ✓ 814ms | ✓ 1293ms | ✓ 789ms | http |
| 142.171.224.229:7890 | ✓ 353ms | ✓ 1905ms | ✓ 1366ms | ✓ 852ms | ✓ 611ms | http |
| 172.212.68.37:3128 | ✓ 310ms | ✓ 1677ms | ✓ 1954ms | 否 | 否 | http |
| 167.71.196.28:8080 | ✓ 805ms | 否 | ✓ 1462ms | ✓ 1135ms | 否 | http |
| 103.172.23.70:1111 | ✓ 1462ms | 否 | 否 | ✓ 1542ms | ✓ 1497ms | http |
| 165.227.152.25:3128 | ✓ 1091ms | 否 | ✓ 1580ms | ✓ 1584ms | ✓ 1336ms | http |
| 195.123.209.48:3128 | ✓ 1105ms | ✓ 1917ms | ✓ 1664ms | ✓ 1978ms | ✓ 1783ms | http |
| 160.250.4.245:1 | ✓ 1519ms | 否 | ✓ 1342ms | ✓ 1437ms | ✓ 1057ms | http |
| 45.136.130.185:8444 | ✓ 1101ms | 否 | ✓ 797ms | ✓ 1237ms | ✓ 1549ms | http |
| 59.11.138.229:3128 | 否 | 否 | ✓ 1531ms | ✓ 1057ms | ✓ 902ms | http |
| 137.184.1.155:3128 | ✓ 392ms | 否 | ✓ 897ms | ✓ 832ms | 否 | http |
| 8.219.97.248:80 | ✓ 1792ms | 否 | ✓ 1668ms | 否 | ✓ 1834ms | http |
| 166.88.55.83:7890 | ✓ 764ms | ✓ 1588ms | ✓ 770ms | ✓ 882ms | ✓ 717ms | http |
| 210.45.70.16:7895 | ✓ 1048ms | ✓ 1316ms | ✓ 1213ms | ✓ 1292ms | ✓ 1035ms | http |
| 106.117.208.101:7890 | ✓ 1012ms | 否 | ✓ 1458ms | 否 | ✓ 1904ms | http |
| 181.41.201.85:3128 | ✓ 1318ms | 否 | ✓ 1812ms | 否 | ✓ 1586ms | http |
| 219.117.204.211:7799 | ✓ 678ms | 否 | ✓ 1051ms | ✓ 1277ms | ✓ 979ms | http |
| 8.217.106.71:8888 | ✓ 717ms | 否 | ✓ 730ms | ✓ 903ms | ✓ 725ms | http |
| 103.183.10.169:3125 | ✓ 1475ms | 否 | ✓ 1854ms | ✓ 1616ms | ✓ 1571ms | http |
| 193.23.200.251:10808 | ✓ 1783ms | 否 | ✓ 1356ms | 否 | ✓ 1860ms | http |
| 62.234.206.73:3128 | ✓ 1279ms | ✓ 1336ms | ✓ 1496ms | 否 | 否 | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1282ms | ✓ 1508ms | ✓ 1037ms | http |
| 210.223.44.230:3128 | ✓ 1548ms | ✓ 1945ms | ✓ 777ms | ✓ 1067ms | ✓ 902ms | http |
| 103.82.23.118:5242 | ✓ 1423ms | 否 | ✓ 1646ms | 否 | ✓ 1766ms | http |
| 103.82.23.118:5171 | ✓ 1453ms | 否 | ✓ 1180ms | 否 | ✓ 1870ms | http |
| 121.230.9.205:1080 | ✓ 1900ms | ✓ 1420ms | ✓ 1171ms | 否 | ✓ 1264ms | http |
| 137.220.151.110:6005 | ✓ 1824ms | 否 | ✓ 1674ms | ✓ 1747ms | ✓ 1485ms | http |
| 137.184.6.37:3128 | ✓ 376ms | ✓ 1658ms | ✓ 1299ms | ✓ 818ms | ✓ 645ms | http |
| 47.101.159.19:8899 | ✓ 948ms | ✓ 1167ms | ✓ 1019ms | ✓ 1225ms | ✓ 943ms | http |
| 121.230.9.19:1080 | ✓ 1201ms | 否 | ✓ 1098ms | 否 | ✓ 1183ms | http |
| 143.244.140.119:3128 | ✓ 1281ms | 否 | ✓ 1755ms | ✓ 1469ms | ✓ 1532ms | http |
| 121.230.8.208:1080 | ✓ 1814ms | ✓ 1798ms | ✓ 1242ms | ✓ 1958ms | ✓ 1299ms | http |
| 213.131.85.29:1976 | ✓ 1899ms | 否 | ✓ 1062ms | ✓ 1720ms | 否 | http |
| 202.141.161.53:30001 | ✓ 1831ms | 否 | ✓ 1335ms | ✓ 1443ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1970ms | 否 | ✓ 1185ms | ✓ 1686ms | 否 | http |
| 218.89.134.230:3333 | 否 | 否 | ✓ 1547ms | ✓ 1857ms | ✓ 1296ms | http |

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
