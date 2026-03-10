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

最后更新：2026-03-10 07:46:19 UTC（2026-03-10 15:46:19 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.47:8443 | ✓ 805ms | ✓ 722ms | ✓ 923ms | ✓ 807ms | ✓ 553ms | http |
| 154.3.236.202:3128 | ✓ 567ms | 否 | ✓ 245ms | ✓ 1130ms | ✓ 1012ms | http |
| 46.183.25.8:443 | ✓ 1065ms | 否 | 否 | ✓ 1568ms | ✓ 1292ms | http |
| 190.212.131.238:3128 | ✓ 910ms | 否 | 否 | ✓ 1568ms | ✓ 1199ms | http |
| 178.236.245.59:3128 | ✓ 927ms | 否 | ✓ 1054ms | 否 | ✓ 1747ms | http |
| 95.3.9.78:3128 | ✓ 1002ms | 否 | 否 | ✓ 1784ms | ✓ 1726ms | http |
| 185.191.236.162:3128 | ✓ 1788ms | ✓ 1680ms | ✓ 1586ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 1823ms | 否 | ✓ 1526ms | ✓ 1602ms | ✓ 1337ms | http |
| 35.225.22.61:80 | ✓ 1564ms | ✓ 1222ms | ✓ 1123ms | 否 | ✓ 712ms | http |
| 168.235.110.63:3128 | ✓ 1434ms | 否 | ✓ 812ms | ✓ 1412ms | ✓ 879ms | http |
| 178.236.245.17:3128 | ✓ 698ms | ✓ 1889ms | ✓ 1371ms | ✓ 1737ms | ✓ 1564ms | http |
| 210.223.44.230:3128 | ✓ 1395ms | 否 | ✓ 1369ms | ✓ 1262ms | 否 | http |
| 152.70.98.46:8888 | ✓ 1387ms | 否 | ✓ 1455ms | ✓ 817ms | ✓ 646ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1619ms | ✓ 1998ms | ✓ 1141ms | http |
| 115.231.181.40:8128 | ✓ 922ms | ✓ 1173ms | ✓ 964ms | ✓ 1198ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1052ms | 否 | ✓ 1328ms | 否 | ✓ 1604ms | http |
| 165.227.5.10:8888 | ✓ 330ms | 否 | 否 | ✓ 1030ms | ✓ 687ms | http |
| 81.70.169.194:80 | ✓ 980ms | ✓ 1409ms | ✓ 1316ms | ✓ 1236ms | 否 | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1002ms | ✓ 1281ms | ✓ 987ms | http |
| 45.136.130.223:8443 | 否 | ✓ 1089ms | ✓ 349ms | 否 | ✓ 569ms | http |
| 95.3.9.78:8080 | ✓ 1297ms | 否 | ✓ 1804ms | ✓ 1718ms | ✓ 1517ms | http |
| 202.155.12.161:443 | ✓ 1924ms | 否 | ✓ 1775ms | 否 | ✓ 1529ms | http |
| 101.43.255.96:80 | ✓ 1077ms | ✓ 1361ms | ✓ 1615ms | ✓ 1566ms | 否 | http |
| 39.104.201.40:7890 | ✓ 1873ms | ✓ 1213ms | 否 | ✓ 1227ms | 否 | http |
| 190.9.109.198:999 | ✓ 913ms | 否 | ✓ 1231ms | ✓ 1639ms | ✓ 1326ms | http |
| 194.213.18.200:443 | ✓ 583ms | ✓ 1822ms | ✓ 596ms | ✓ 1314ms | 否 | http |
| 193.168.173.136:443 | ✓ 881ms | ✓ 1770ms | ✓ 1256ms | 否 | 否 | http |
| 14.225.222.213:7890 | ✓ 1850ms | 否 | ✓ 1357ms | 否 | ✓ 920ms | http |
| 120.92.212.16:7890 | ✓ 1097ms | 否 | 否 | ✓ 1342ms | ✓ 994ms | http |
| 14.225.212.37:7890 | ✓ 1410ms | 否 | 否 | ✓ 1755ms | ✓ 1088ms | http |
| 94.176.3.43:7443 | ✓ 1048ms | 否 | ✓ 1250ms | ✓ 1702ms | 否 | http |
| 183.249.5.117:22222 | 否 | 否 | ✓ 802ms | ✓ 1189ms | ✓ 1855ms | http |
| 106.14.203.63:3333 | ✓ 1720ms | ✓ 1754ms | ✓ 1694ms | 否 | ✓ 1980ms | http |
| 34.101.184.164:3128 | ✓ 1936ms | 否 | ✓ 1677ms | ✓ 1466ms | ✓ 1287ms | http |
| 205.209.118.30:3138 | ✓ 386ms | ✓ 1372ms | ✓ 1716ms | 否 | 否 | http |
| 47.77.193.180:1080 | ✓ 446ms | ✓ 844ms | ✓ 515ms | ✓ 782ms | ✓ 546ms | http |
| 45.136.198.40:3128 | ✓ 1264ms | 否 | ✓ 1430ms | 否 | ✓ 1713ms | http |
| 91.107.141.42:8081 | ✓ 1202ms | 否 | 否 | ✓ 1963ms | ✓ 1772ms | http |
| 150.107.140.238:3128 | ✓ 1646ms | 否 | ✓ 1228ms | ✓ 1623ms | ✓ 1210ms | http |
| 59.46.216.131:30001 | ✓ 1142ms | ✓ 1806ms | ✓ 1210ms | ✓ 1393ms | 否 | http |
| 103.113.70.189:1081 | ✓ 970ms | 否 | 否 | ✓ 1441ms | ✓ 1946ms | http |
| 62.113.119.14:8080 | ✓ 1065ms | 否 | ✓ 822ms | ✓ 1561ms | ✓ 1191ms | http |
| 45.140.147.155:1081 | ✓ 1057ms | 否 | ✓ 1090ms | ✓ 1641ms | ✓ 1173ms | http |
| 116.80.82.224:3172 | ✓ 1564ms | 否 | ✓ 1553ms | ✓ 1907ms | 否 | http |
| 8.219.97.248:80 | ✓ 1720ms | 否 | ✓ 1526ms | ✓ 1706ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1037ms | 否 | ✓ 1153ms | 否 | ✓ 1348ms | http |
| 121.237.181.137:8888 | 否 | 否 | ✓ 1921ms | ✓ 1809ms | ✓ 856ms | http |
| 45.140.147.82:1081 | ✓ 676ms | 否 | ✓ 845ms | 否 | ✓ 1320ms | http |
| 120.238.159.229:22222 | ✓ 915ms | ✓ 1252ms | ✓ 1521ms | ✓ 1205ms | ✓ 920ms | http |
| 113.59.32.161:22222 | ✓ 1103ms | ✓ 1382ms | ✓ 1128ms | 否 | 否 | http |
| 117.159.239.58:22222 | ✓ 856ms | ✓ 1072ms | ✓ 861ms | ✓ 1121ms | ✓ 861ms | http |
| 152.42.213.210:8080 | ✓ 895ms | 否 | ✓ 1096ms | ✓ 1086ms | ✓ 849ms | http |
| 120.238.159.230:22222 | 否 | 否 | ✓ 1199ms | ✓ 1133ms | ✓ 933ms | http |
| 113.59.32.162:22222 | ✓ 1246ms | ✓ 1407ms | ✓ 1160ms | ✓ 1327ms | ✓ 959ms | http |
| 183.249.5.111:22222 | ✓ 876ms | ✓ 1298ms | ✓ 884ms | ✓ 1237ms | ✓ 840ms | http |
| 113.59.32.142:22222 | ✓ 1146ms | ✓ 1412ms | ✓ 1205ms | ✓ 1372ms | ✓ 1077ms | http |
| 183.249.5.109:22222 | 否 | 否 | ✓ 1552ms | ✓ 1204ms | ✓ 898ms | http |
| 120.238.159.228:22222 | ✓ 938ms | ✓ 1260ms | ✓ 997ms | ✓ 1123ms | ✓ 913ms | http |
| 120.238.159.227:22222 | ✓ 960ms | 否 | 否 | ✓ 1166ms | ✓ 887ms | http |
| 120.240.29.173:22222 | ✓ 993ms | ✓ 1213ms | ✓ 1033ms | ✓ 1264ms | ✓ 926ms | http |
| 120.240.29.51:22222 | 否 | ✓ 1307ms | ✓ 1133ms | ✓ 1178ms | ✓ 906ms | http |
| 103.39.51.190:8080 | ✓ 1351ms | 否 | 否 | ✓ 1768ms | ✓ 1642ms | http |
| 61.52.131.172:8443 | ✓ 856ms | ✓ 1148ms | ✓ 990ms | ✓ 1176ms | ✓ 944ms | http |
| 152.42.213.210:80 | ✓ 915ms | 否 | 否 | ✓ 1639ms | ✓ 1127ms | http |
| 222.184.48.252:22222 | ✓ 1971ms | ✓ 1701ms | ✓ 1003ms | ✓ 1228ms | ✓ 948ms | http |
| 117.159.239.49:22222 | ✓ 885ms | ✓ 1098ms | ✓ 886ms | ✓ 1233ms | ✓ 981ms | http |
| 117.159.239.50:22222 | 否 | 否 | ✓ 838ms | ✓ 1092ms | ✓ 874ms | http |
| 116.80.49.169:3172 | ✓ 1786ms | 否 | ✓ 1551ms | ✓ 1870ms | 否 | http |

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
