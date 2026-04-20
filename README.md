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

最后更新：2026-04-20 16:11:45 UTC（2026-04-21 00:11:45 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 178ms | 否 | ✓ 772ms | 否 | ✓ 879ms | http |
| 46.101.95.183:8888 | 否 | 否 | ✓ 640ms | ✓ 1434ms | ✓ 1156ms | http |
| 152.32.132.190:7890 | ✓ 1026ms | 否 | 否 | ✓ 1078ms | ✓ 1035ms | http |
| 152.42.208.139:8118 | ✓ 1059ms | 否 | ✓ 952ms | ✓ 1185ms | 否 | http |
| 185.138.116.150:8080 | ✓ 1050ms | 否 | ✓ 1559ms | 否 | ✓ 1329ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1505ms | ✓ 1392ms | ✓ 1433ms | ✓ 1110ms | http |
| 91.99.15.45:2095 | 否 | 否 | ✓ 1348ms | ✓ 1750ms | ✓ 1758ms | http |
| 81.30.156.115:8080 | 否 | 否 | ✓ 1321ms | ✓ 1701ms | ✓ 1713ms | http |
| 168.144.75.9:3128 | ✓ 1887ms | 否 | ✓ 1933ms | 否 | ✓ 1643ms | http |
| 195.26.224.49:3128 | 否 | 否 | ✓ 635ms | ✓ 1505ms | ✓ 1311ms | http |
| 149.51.42.10:8080 | ✓ 403ms | ✓ 1299ms | 否 | ✓ 1256ms | 否 | http |
| 159.89.191.221:3128 | ✓ 225ms | 否 | 否 | ✓ 1157ms | ✓ 1263ms | http |
| 1.231.81.166:3128 | ✓ 838ms | ✓ 1874ms | ✓ 1023ms | ✓ 1005ms | ✓ 825ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1304ms | ✓ 964ms | ✓ 1317ms | 否 | http |
| 188.246.224.49:7890 | ✓ 536ms | 否 | ✓ 1144ms | ✓ 1450ms | ✓ 1442ms | http |
| 177.93.132.244:3128 | ✓ 1160ms | 否 | ✓ 737ms | 否 | ✓ 1678ms | http |
| 130.61.30.221:8080 | ✓ 590ms | 否 | ✓ 1638ms | 否 | ✓ 1727ms | http |
| 218.108.131.186:17890 | ✓ 1843ms | 否 | ✓ 1794ms | 否 | ✓ 1824ms | http |
| 20.210.39.153:8561 | ✓ 1587ms | 否 | ✓ 1985ms | 否 | ✓ 1698ms | http |
| 20.78.118.91:8561 | ✓ 1587ms | 否 | ✓ 1991ms | 否 | ✓ 1700ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1372ms | ✓ 1068ms | 否 | ✓ 1840ms | http |
| 162.240.154.26:3128 | ✓ 1595ms | 否 | ✓ 1661ms | 否 | ✓ 1543ms | http |
| 192.3.248.190:8014 | ✓ 604ms | 否 | ✓ 1230ms | ✓ 1454ms | ✓ 1206ms | http |
| 27.71.24.102:3128 | ✓ 1632ms | 否 | ✓ 1357ms | ✓ 1235ms | ✓ 993ms | http |
| 103.113.70.189:1081 | ✓ 1338ms | ✓ 1337ms | ✓ 1242ms | ✓ 1483ms | 否 | http |
| 106.10.55.212:1121 | ✓ 876ms | 否 | ✓ 1049ms | ✓ 1222ms | ✓ 1261ms | http |
| 212.58.132.5:8888 | ✓ 1168ms | 否 | ✓ 1725ms | ✓ 1560ms | ✓ 1260ms | http |
| 14.143.222.113:57788 | ✓ 1032ms | 否 | ✓ 1123ms | ✓ 1458ms | ✓ 1336ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1518ms | ✓ 1825ms | ✓ 1469ms | ✓ 1186ms | http |
| 45.153.231.229:8080 | 否 | 否 | ✓ 1961ms | ✓ 1918ms | ✓ 1510ms | http |
| 45.12.151.226:2829 | 否 | ✓ 1708ms | ✓ 1298ms | 否 | ✓ 1399ms | http |
| 168.222.254.136:8888 | ✓ 1705ms | 否 | ✓ 1411ms | ✓ 1997ms | ✓ 1553ms | http |
| 83.219.250.8:62920 | ✓ 1062ms | 否 | ✓ 1304ms | 否 | ✓ 1775ms | http |
| 34.71.229.255:3128 | ✓ 234ms | ✓ 1387ms | ✓ 1657ms | ✓ 1164ms | ✓ 993ms | http |
| 157.230.178.216:8088 | ✓ 958ms | ✓ 1913ms | ✓ 1091ms | ✓ 1480ms | ✓ 1409ms | http |
| 178.213.25.221:7890 | ✓ 1061ms | 否 | ✓ 1276ms | ✓ 1784ms | ✓ 1445ms | http |
| 85.190.99.143:443 | ✓ 1292ms | 否 | ✓ 947ms | 否 | ✓ 1825ms | http |
| 84.47.150.125:1080 | ✓ 1035ms | 否 | ✓ 1832ms | 否 | ✓ 1617ms | http |
| 137.59.47.73:3128 | ✓ 1558ms | 否 | 否 | ✓ 1986ms | ✓ 1598ms | http |
| 193.23.194.147:3128 | ✓ 1261ms | 否 | ✓ 1719ms | 否 | ✓ 1549ms | http |
| 149.51.42.10:3128 | ✓ 1568ms | ✓ 1625ms | 否 | ✓ 1657ms | 否 | http |
| 103.113.70.189:1082 | ✓ 669ms | ✓ 1113ms | ✓ 722ms | ✓ 1314ms | ✓ 1198ms | http |
| 47.84.73.61:1080 | ✓ 1581ms | ✓ 1885ms | ✓ 1018ms | ✓ 1191ms | ✓ 989ms | http |
| 202.141.161.53:10808 | ✓ 1169ms | 否 | 否 | ✓ 1396ms | ✓ 1962ms | http |
| 78.11.96.22:8888 | ✓ 926ms | 否 | ✓ 1062ms | ✓ 1513ms | ✓ 1354ms | http |
| 45.140.147.155:1081 | ✓ 1505ms | 否 | ✓ 1423ms | 否 | ✓ 1411ms | http |
| 103.85.113.66:9999 | ✓ 1166ms | ✓ 1904ms | ✓ 1364ms | ✓ 1859ms | 否 | http |
| 117.236.124.166:3128 | ✓ 1622ms | 否 | ✓ 1077ms | ✓ 1848ms | ✓ 1480ms | http |
| 208.87.243.199:7878 | 否 | 否 | ✓ 1107ms | ✓ 1006ms | ✓ 1186ms | http |
| 138.124.99.216:8888 | ✓ 1718ms | 否 | ✓ 868ms | 否 | ✓ 1422ms | http |
| 46.175.148.17:2040 | ✓ 605ms | ✓ 1294ms | 否 | ✓ 1935ms | ✓ 1549ms | http |
| 38.180.2.107:3128 | ✓ 784ms | ✓ 1693ms | ✓ 1992ms | 否 | ✓ 1936ms | http |
| 62.113.119.14:8080 | ✓ 637ms | ✓ 1801ms | ✓ 1258ms | 否 | ✓ 1809ms | http |
| 101.32.243.189:80 | 否 | 否 | ✓ 1588ms | ✓ 1719ms | ✓ 1428ms | http |
| 164.92.166.127:3128 | ✓ 452ms | ✓ 1871ms | ✓ 1908ms | ✓ 1717ms | ✓ 1349ms | http |
| 160.250.134.143:3128 | 否 | 否 | ✓ 1051ms | ✓ 1431ms | ✓ 1079ms | http |
| 147.45.166.46:3128 | ✓ 1813ms | 否 | 否 | ✓ 1858ms | ✓ 1797ms | http |
| 45.93.30.241:6005 | 否 | 否 | ✓ 1863ms | ✓ 1716ms | ✓ 1580ms | http |
| 45.93.29.147:6005 | ✓ 1941ms | 否 | ✓ 1823ms | ✓ 1435ms | 否 | http |
| 103.93.93.221:8181 | ✓ 1854ms | 否 | ✓ 1612ms | ✓ 1637ms | ✓ 1537ms | http |
| 91.193.240.157:9877 | ✓ 863ms | 否 | ✓ 817ms | ✓ 1912ms | ✓ 1781ms | http |
| 193.181.35.119:8118 | ✓ 860ms | 否 | ✓ 1418ms | 否 | ✓ 1752ms | http |
| 94.131.118.129:1081 | ✓ 940ms | ✓ 1698ms | ✓ 1317ms | 否 | ✓ 1987ms | http |
| 42.200.76.16:3888 | 否 | 否 | ✓ 878ms | ✓ 1055ms | ✓ 844ms | http |
| 139.227.17.70:17890 | 否 | 否 | ✓ 925ms | ✓ 1235ms | ✓ 982ms | http |
| 45.140.147.82:1081 | ✓ 597ms | ✓ 1544ms | ✓ 899ms | ✓ 1654ms | ✓ 1224ms | http |
| 113.176.92.71:3128 | ✓ 1301ms | ✓ 1444ms | ✓ 1460ms | ✓ 1332ms | ✓ 1027ms | http |
| 43.132.188.134:443 | ✓ 1006ms | ✓ 1365ms | 否 | ✓ 1784ms | 否 | http |
| 218.153.163.186:8508 | 否 | 否 | ✓ 1352ms | ✓ 1653ms | ✓ 993ms | http |
| 103.177.20.21:8181 | ✓ 1639ms | 否 | ✓ 1599ms | ✓ 1715ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1453ms | 否 | ✓ 1275ms | 否 | ✓ 1296ms | http |
| 103.247.23.37:1111 | ✓ 1934ms | 否 | ✓ 1963ms | 否 | ✓ 1698ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1487ms | ✓ 700ms | 否 | ✓ 812ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1669ms | 否 | ✓ 1618ms | ✓ 1116ms | http |
| 186.96.15.86:8080 | ✓ 1576ms | 否 | 否 | ✓ 1599ms | ✓ 1087ms | http |
| 46.39.105.157:8080 | ✓ 661ms | 否 | ✓ 1019ms | ✓ 1597ms | ✓ 1412ms | http |
| 223.84.151.86:30005 | ✓ 1385ms | 否 | ✓ 1579ms | ✓ 1689ms | ✓ 1417ms | http |
| 61.52.131.172:8443 | ✓ 1011ms | ✓ 1350ms | ✓ 993ms | ✓ 1371ms | ✓ 1047ms | http |
| 8.140.104.98:3128 | ✓ 1077ms | ✓ 1353ms | ✓ 1056ms | ✓ 1386ms | ✓ 1163ms | http |
| 20.27.13.35:8561 | ✓ 1525ms | ✓ 1700ms | ✓ 1911ms | 否 | 否 | http |

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
