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

最后更新：2026-04-17 15:49:46 UTC（2026-04-17 23:49:46 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:8080 | ✓ 871ms | ✓ 1718ms | 否 | ✓ 1645ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1261ms | ✓ 962ms | ✓ 916ms | ✓ 909ms | ✓ 662ms | http |
| 113.160.132.26:8080 | ✓ 1671ms | 否 | ✓ 941ms | ✓ 1257ms | ✓ 948ms | http |
| 185.138.116.150:8080 | ✓ 1118ms | 否 | 否 | ✓ 1919ms | ✓ 1528ms | http |
| 20.127.128.70:8080 | ✓ 1512ms | 否 | ✓ 1604ms | 否 | ✓ 1988ms | http |
| 78.11.96.22:8888 | ✓ 1067ms | 否 | 否 | ✓ 1828ms | ✓ 1716ms | http |
| 103.113.70.189:1082 | ✓ 734ms | ✓ 1874ms | ✓ 485ms | ✓ 1430ms | ✓ 1009ms | http |
| 218.108.131.186:17890 | ✓ 849ms | ✓ 998ms | ✓ 1002ms | ✓ 1066ms | ✓ 840ms | http |
| 103.113.70.189:1081 | ✓ 687ms | 否 | ✓ 415ms | 否 | ✓ 1064ms | http |
| 223.84.151.86:30005 | ✓ 1368ms | ✓ 1177ms | ✓ 1021ms | ✓ 1285ms | ✓ 1178ms | http |
| 36.141.21.200:7890 | ✓ 1596ms | 否 | ✓ 867ms | ✓ 1170ms | ✓ 929ms | http |
| 188.246.224.49:7890 | ✓ 733ms | 否 | ✓ 864ms | 否 | ✓ 1906ms | http |
| 91.233.223.147:3128 | ✓ 1258ms | 否 | ✓ 1210ms | 否 | ✓ 1718ms | http |
| 168.144.75.9:3128 | ✓ 1211ms | 否 | ✓ 1555ms | ✓ 1982ms | ✓ 1482ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1259ms | ✓ 1223ms | 否 | ✓ 1990ms | http |
| 51.145.178.158:3128 | ✓ 603ms | ✓ 1699ms | ✓ 615ms | ✓ 1712ms | 否 | http |
| 208.87.243.199:7878 | ✓ 1301ms | ✓ 1700ms | 否 | ✓ 792ms | ✓ 500ms | http |
| 159.89.191.221:3128 | ✓ 399ms | 否 | ✓ 1189ms | ✓ 1280ms | ✓ 1095ms | http |
| 27.71.24.102:3128 | ✓ 1982ms | 否 | ✓ 1420ms | ✓ 1060ms | ✓ 1326ms | http |
| 130.61.30.221:8080 | ✓ 865ms | 否 | ✓ 1993ms | 否 | ✓ 1887ms | http |
| 89.35.119.147:3128 | ✓ 918ms | 否 | ✓ 1855ms | 否 | ✓ 1983ms | http |
| 70.61.188.34:3128 | ✓ 894ms | 否 | 否 | ✓ 1454ms | ✓ 1324ms | http |
| 103.138.70.165:3129 | ✓ 1290ms | 否 | 否 | ✓ 1404ms | ✓ 1317ms | http |
| 120.92.108.86:7890 | ✓ 1277ms | 否 | ✓ 1443ms | ✓ 1920ms | ✓ 1365ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1884ms | ✓ 1156ms | ✓ 1202ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1612ms | ✓ 1328ms | ✓ 1931ms | http |
| 47.93.216.160:1081 | ✓ 856ms | ✓ 1184ms | ✓ 853ms | ✓ 1231ms | ✓ 992ms | http |
| 210.223.44.230:3128 | ✓ 1947ms | ✓ 1545ms | 否 | 否 | ✓ 1236ms | http |
| 190.12.150.244:999 | ✓ 1662ms | ✓ 1874ms | ✓ 1052ms | 否 | ✓ 1525ms | http |
| 177.93.132.244:3128 | ✓ 1105ms | 否 | ✓ 1130ms | ✓ 1932ms | ✓ 1473ms | http |
| 43.132.188.134:443 | 否 | ✓ 1545ms | ✓ 1855ms | 否 | ✓ 1051ms | http |
| 35.225.22.61:80 | ✓ 516ms | ✓ 1395ms | ✓ 1185ms | ✓ 1240ms | ✓ 948ms | http |
| 84.47.150.125:1080 | ✓ 1402ms | 否 | ✓ 1862ms | 否 | ✓ 1448ms | http |
| 180.125.216.109:8118 | ✓ 1546ms | 否 | 否 | ✓ 1160ms | ✓ 1930ms | http |
| 149.51.42.10:3128 | ✓ 662ms | ✓ 1975ms | 否 | ✓ 1692ms | 否 | http |
| 117.236.124.166:3128 | ✓ 1750ms | 否 | ✓ 1131ms | ✓ 1938ms | 否 | http |
| 158.160.215.167:8124 | ✓ 1256ms | 否 | ✓ 1026ms | 否 | ✓ 1964ms | http |
| 139.227.17.70:17890 | 否 | ✓ 1128ms | ✓ 863ms | ✓ 1134ms | ✓ 892ms | http |
| 34.101.184.164:3128 | ✓ 1673ms | 否 | ✓ 1653ms | ✓ 1522ms | ✓ 1617ms | http |
| 42.101.8.101:8888 | ✓ 1735ms | ✓ 1350ms | 否 | ✓ 1752ms | ✓ 1079ms | http |
| 8.140.104.98:3128 | 否 | 否 | ✓ 1036ms | ✓ 1303ms | ✓ 1037ms | http |
| 101.32.244.83:8080 | ✓ 913ms | ✓ 1524ms | ✓ 919ms | ✓ 1208ms | ✓ 1215ms | http |
| 121.43.196.210:8222 | ✓ 956ms | ✓ 1044ms | ✓ 810ms | ✓ 1062ms | ✓ 909ms | http |
| 121.43.196.213:8222 | ✓ 940ms | ✓ 989ms | ✓ 848ms | ✓ 1078ms | ✓ 933ms | http |
| 114.55.226.123:10086 | ✓ 1071ms | ✓ 1376ms | ✓ 996ms | ✓ 1254ms | ✓ 1060ms | http |
| 47.100.2.5:2020 | ✓ 1585ms | ✓ 946ms | ✓ 1135ms | 否 | ✓ 1933ms | http |
| 138.124.99.216:8888 | ✓ 1318ms | 否 | ✓ 1927ms | 否 | ✓ 1816ms | http |
| 12.89.176.82:3128 | ✓ 487ms | ✓ 1076ms | ✓ 998ms | ✓ 1654ms | ✓ 1062ms | http |
| 116.58.161.203:26021 | ✓ 944ms | 否 | ✓ 1441ms | ✓ 1368ms | ✓ 1101ms | http |
| 212.58.132.5:8888 | ✓ 1830ms | 否 | ✓ 1587ms | ✓ 1485ms | ✓ 1219ms | http |
| 158.160.215.167:8127 | ✓ 1012ms | 否 | ✓ 985ms | ✓ 1981ms | ✓ 1709ms | http |
| 121.230.8.136:1080 | 否 | ✓ 1528ms | ✓ 1316ms | ✓ 1593ms | ✓ 1274ms | http |
| 117.122.240.82:3338 | ✓ 1296ms | ✓ 1616ms | ✓ 1406ms | ✓ 1292ms | ✓ 908ms | http |
| 133.18.123.225:26021 | ✓ 1353ms | 否 | ✓ 1421ms | ✓ 1025ms | ✓ 1427ms | http |
| 62.113.119.14:8080 | ✓ 1367ms | 否 | ✓ 812ms | ✓ 1689ms | ✓ 1537ms | http |
| 211.95.152.50:45046 | 否 | ✓ 1291ms | ✓ 1030ms | ✓ 1182ms | ✓ 944ms | http |
| 217.76.245.80:999 | ✓ 1940ms | 否 | ✓ 1423ms | ✓ 1706ms | 否 | http |
| 108.131.109.106:44593 | ✓ 1946ms | 否 | ✓ 971ms | 否 | ✓ 1841ms | http |
| 202.141.161.53:10808 | ✓ 1903ms | ✓ 1101ms | ✓ 1741ms | 否 | ✓ 1023ms | http |
| 103.85.113.66:9999 | ✓ 1508ms | ✓ 1716ms | ✓ 1918ms | 否 | 否 | http |
| 45.140.147.155:1082 | 否 | 否 | ✓ 607ms | ✓ 1625ms | ✓ 1332ms | http |
| 59.46.216.131:30001 | ✓ 1681ms | ✓ 1777ms | ✓ 1639ms | ✓ 1594ms | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1171ms | ✓ 974ms | ✓ 1366ms | 否 | http |
| 101.132.61.121:8888 | ✓ 1252ms | ✓ 1212ms | ✓ 1260ms | 否 | ✓ 1203ms | http |
| 61.52.131.172:8443 | ✓ 1535ms | 否 | 否 | ✓ 1226ms | ✓ 905ms | http |

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
