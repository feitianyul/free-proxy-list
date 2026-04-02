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

最后更新：2026-04-02 12:46:35 UTC（2026-04-02 20:46:35 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 692ms | ✓ 938ms | ✓ 826ms | ✓ 814ms | ✓ 657ms | http |
| 147.161.210.140:8800 | ✓ 680ms | 否 | ✓ 965ms | ✓ 1309ms | ✓ 899ms | http |
| 165.232.146.249:3128 | 否 | 否 | ✓ 945ms | ✓ 709ms | ✓ 542ms | http |
| 1.231.81.166:3128 | ✓ 780ms | ✓ 1771ms | ✓ 1860ms | ✓ 1233ms | ✓ 975ms | http |
| 203.80.138.81:50000 | 否 | ✓ 1619ms | ✓ 1340ms | ✓ 1085ms | ✓ 947ms | http |
| 167.103.115.102:8800 | ✓ 1440ms | 否 | ✓ 1047ms | ✓ 1043ms | ✓ 1053ms | http |
| 42.96.16.158:1311 | ✓ 1960ms | 否 | ✓ 1030ms | ✓ 1258ms | ✓ 976ms | http |
| 95.213.217.168:52004 | ✓ 1361ms | ✓ 1929ms | ✓ 1180ms | ✓ 1857ms | ✓ 1282ms | http |
| 113.160.132.26:8080 | ✓ 1363ms | 否 | ✓ 1081ms | ✓ 1294ms | ✓ 1331ms | http |
| 208.87.243.199:7878 | ✓ 579ms | 否 | ✓ 1408ms | 否 | ✓ 805ms | http |
| 167.103.34.108:8800 | ✓ 1742ms | 否 | ✓ 1628ms | 否 | ✓ 1530ms | http |
| 45.167.124.52:8080 | ✓ 1299ms | 否 | ✓ 1365ms | 否 | ✓ 1681ms | http |
| 35.225.22.61:80 | ✓ 890ms | 否 | 否 | ✓ 1195ms | ✓ 1142ms | http |
| 167.103.144.127:8800 | ✓ 1190ms | 否 | ✓ 1155ms | ✓ 1202ms | ✓ 1091ms | http |
| 167.160.191.204:6005 | ✓ 641ms | 否 | ✓ 1429ms | 否 | ✓ 1705ms | http |
| 212.58.132.5:8888 | ✓ 1122ms | 否 | ✓ 1487ms | ✓ 1513ms | ✓ 1511ms | http |
| 167.103.31.122:8800 | ✓ 1736ms | 否 | ✓ 1938ms | ✓ 1804ms | 否 | http |
| 217.76.245.80:999 | ✓ 985ms | ✓ 1527ms | ✓ 1154ms | 否 | ✓ 1192ms | http |
| 45.167.125.21:999 | ✓ 762ms | ✓ 1916ms | ✓ 1197ms | ✓ 1769ms | ✓ 1451ms | http |
| 119.28.156.42:3128 | ✓ 1479ms | 否 | ✓ 655ms | ✓ 843ms | ✓ 715ms | http |
| 116.80.96.103:3172 | ✓ 1533ms | 否 | ✓ 1500ms | 否 | ✓ 1609ms | http |
| 128.199.116.219:9090 | ✓ 1470ms | 否 | ✓ 1086ms | ✓ 1092ms | ✓ 937ms | http |
| 159.223.71.162:8080 | ✓ 1471ms | 否 | ✓ 1185ms | ✓ 1073ms | ✓ 862ms | http |
| 159.223.71.162:443 | ✓ 1470ms | 否 | ✓ 1198ms | ✓ 1070ms | ✓ 854ms | http |
| 128.199.121.61:9090 | ✓ 1520ms | 否 | ✓ 1098ms | ✓ 1110ms | 否 | http |
| 147.161.239.240:8800 | ✓ 1187ms | 否 | ✓ 1451ms | ✓ 1750ms | ✓ 1564ms | http |
| 101.43.127.100:8877 | ✓ 1235ms | 否 | ✓ 854ms | ✓ 1149ms | 否 | http |
| 160.250.5.22:1 | ✓ 1764ms | 否 | ✓ 1882ms | ✓ 1729ms | ✓ 1637ms | http |
| 177.234.217.88:999 | ✓ 1932ms | 否 | ✓ 1746ms | ✓ 1874ms | ✓ 1608ms | http |
| 128.199.113.85:9090 | 否 | 否 | ✓ 1226ms | ✓ 1230ms | ✓ 897ms | http |
| 146.190.80.158:9090 | ✓ 758ms | 否 | ✓ 712ms | ✓ 1046ms | ✓ 842ms | http |
| 104.248.243.244:3128 | ✓ 595ms | ✓ 1886ms | 否 | ✓ 1965ms | ✓ 1484ms | http |
| 45.12.151.226:2829 | ✓ 1919ms | 否 | ✓ 1470ms | 否 | ✓ 1474ms | http |
| 150.241.71.15:1080 | ✓ 685ms | 否 | ✓ 1031ms | ✓ 1967ms | ✓ 1138ms | http |
| 115.231.181.40:8128 | ✓ 1911ms | ✓ 1267ms | ✓ 955ms | 否 | 否 | http |
| 174.140.109.250:3128 | ✓ 596ms | 否 | ✓ 575ms | ✓ 1804ms | ✓ 1109ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 941ms | ✓ 1277ms | ✓ 989ms | http |
| 106.10.55.212:1121 | ✓ 911ms | ✓ 1945ms | 否 | ✓ 1207ms | ✓ 1309ms | http |
| 106.117.208.101:7890 | ✓ 1046ms | ✓ 1582ms | 否 | ✓ 1267ms | ✓ 1989ms | http |
| 116.80.49.162:3172 | ✓ 1553ms | 否 | 否 | ✓ 1917ms | ✓ 1718ms | http |
| 101.32.244.83:8080 | ✓ 1031ms | 否 | ✓ 947ms | ✓ 1411ms | ✓ 1204ms | http |
| 121.43.196.213:8222 | ✓ 939ms | ✓ 1059ms | ✓ 930ms | ✓ 1081ms | ✓ 921ms | http |
| 121.43.196.210:8222 | 否 | ✓ 1089ms | ✓ 868ms | ✓ 1072ms | ✓ 909ms | http |
| 114.55.226.123:10086 | ✓ 1121ms | ✓ 1396ms | ✓ 1137ms | ✓ 1299ms | ✓ 1064ms | http |
| 150.107.140.238:3128 | ✓ 935ms | 否 | ✓ 1229ms | ✓ 1328ms | 否 | http |
| 45.136.130.173:8448 | ✓ 1799ms | ✓ 1654ms | ✓ 812ms | 否 | 否 | http |
| 38.145.218.82:8443 | ✓ 1132ms | 否 | ✓ 137ms | ✓ 1445ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1073ms | ✓ 1821ms | ✓ 1928ms | 否 | 否 | http |
| 38.145.208.211:8443 | ✓ 698ms | 否 | ✓ 997ms | ✓ 660ms | ✓ 575ms | http |
| 128.199.254.13:9090 | ✓ 1752ms | 否 | 否 | ✓ 1223ms | ✓ 882ms | http |
| 150.249.255.91:3128 | ✓ 635ms | 否 | ✓ 529ms | ✓ 809ms | ✓ 661ms | http |
| 104.248.151.93:9090 | ✓ 805ms | 否 | ✓ 1075ms | ✓ 1070ms | ✓ 844ms | http |
| 38.34.179.152:8446 | 否 | ✓ 1697ms | ✓ 1962ms | ✓ 860ms | ✓ 838ms | http |
| 195.123.209.48:3128 | ✓ 820ms | 否 | ✓ 1765ms | 否 | ✓ 1841ms | http |
| 1.234.153.14:80 | ✓ 1429ms | ✓ 988ms | ✓ 751ms | ✓ 861ms | ✓ 652ms | http |
| 103.82.23.118:5234 | ✓ 1769ms | 否 | ✓ 1095ms | 否 | ✓ 1723ms | http |
| 210.223.44.230:3128 | ✓ 732ms | 否 | ✓ 1922ms | ✓ 915ms | ✓ 871ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1924ms | ✓ 1494ms | ✓ 1596ms | http |
| 181.78.44.63:999 | ✓ 852ms | 否 | ✓ 1110ms | ✓ 1756ms | ✓ 1462ms | http |
| 116.80.65.77:3172 | 否 | 否 | ✓ 1801ms | ✓ 1991ms | ✓ 1692ms | http |
| 120.92.212.16:8890 | ✓ 953ms | ✓ 1259ms | 否 | 否 | ✓ 1681ms | http |
| 101.230.73.57:29999 | ✓ 807ms | ✓ 969ms | ✓ 820ms | 否 | ✓ 1053ms | http |
| 82.114.228.67:1080 | ✓ 822ms | 否 | ✓ 1408ms | 否 | ✓ 1275ms | http |
| 128.199.114.189:9090 | ✓ 1510ms | 否 | ✓ 1014ms | ✓ 1128ms | ✓ 835ms | http |
| 116.80.65.83:3172 | ✓ 1937ms | 否 | ✓ 1490ms | ✓ 1997ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1940ms | 否 | 否 | ✓ 1144ms | ✓ 1059ms | http |
| 209.126.84.232:8888 | ✓ 1738ms | 否 | ✓ 1672ms | ✓ 1724ms | ✓ 1288ms | http |
| 207.254.71.62:8088 | ✓ 761ms | 否 | ✓ 860ms | ✓ 1746ms | ✓ 1577ms | http |
| 165.225.72.38:11376 | 否 | 否 | ✓ 1040ms | ✓ 1877ms | ✓ 1482ms | http |
| 186.96.111.214:999 | ✓ 1009ms | ✓ 1740ms | ✓ 1542ms | ✓ 1693ms | ✓ 1571ms | http |
| 61.52.131.172:8443 | ✓ 914ms | ✓ 1097ms | ✓ 887ms | ✓ 1142ms | ✓ 916ms | http |
| 121.230.8.250:1080 | ✓ 1815ms | 否 | ✓ 1999ms | ✓ 1764ms | ✓ 1971ms | http |
| 200.125.171.254:999 | 否 | ✓ 1709ms | ✓ 1450ms | ✓ 1672ms | 否 | http |
| 46.101.190.71:3128 | ✓ 614ms | 否 | ✓ 1698ms | ✓ 1862ms | ✓ 1488ms | http |
| 141.136.63.126:8080 | 否 | 否 | ✓ 1892ms | ✓ 1960ms | ✓ 1697ms | http |

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
