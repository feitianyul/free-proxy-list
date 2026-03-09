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

最后更新：2026-03-09 20:51:09 UTC（2026-03-10 04:51:09 UTC+8）

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
| 101.47.73.135:3128 | 否 | 否 | ✓ 1871ms | ✓ 1471ms | ✓ 999ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1268ms | ✓ 1146ms | ✓ 941ms | http |
| 162.240.154.26:3128 | ✓ 1490ms | ✓ 1412ms | ✓ 1254ms | ✓ 1588ms | ✓ 1107ms | http |
| 115.231.181.40:8128 | ✓ 715ms | 否 | 否 | ✓ 1012ms | ✓ 794ms | http |
| 113.177.131.2:3128 | 否 | ✓ 1937ms | 否 | ✓ 1358ms | ✓ 1127ms | http |
| 38.210.179.106:999 | ✓ 1488ms | ✓ 1458ms | ✓ 1328ms | ✓ 1763ms | ✓ 1491ms | http |
| 165.227.5.10:8888 | ✓ 47ms | ✓ 1336ms | ✓ 305ms | 否 | ✓ 486ms | http |
| 101.43.255.96:80 | 否 | ✓ 1572ms | ✓ 812ms | ✓ 1591ms | ✓ 828ms | http |
| 121.237.181.137:8888 | ✓ 755ms | ✓ 927ms | ✓ 782ms | ✓ 905ms | ✓ 749ms | http |
| 81.70.169.194:80 | ✓ 1780ms | ✓ 1077ms | ✓ 792ms | ✓ 1934ms | ✓ 1775ms | http |
| 190.9.109.198:999 | ✓ 1214ms | ✓ 1536ms | ✓ 1174ms | ✓ 1459ms | ✓ 1225ms | http |
| 190.9.109.207:999 | ✓ 1221ms | ✓ 1574ms | ✓ 1245ms | ✓ 1372ms | ✓ 1247ms | http |
| 193.168.173.136:443 | ✓ 1216ms | 否 | ✓ 1531ms | ✓ 1894ms | 否 | http |
| 1.231.81.166:3128 | 否 | ✓ 1856ms | ✓ 1530ms | ✓ 1592ms | ✓ 1523ms | http |
| 46.183.25.8:443 | ✓ 834ms | 否 | ✓ 1803ms | ✓ 1440ms | 否 | http |
| 202.155.12.161:443 | ✓ 1296ms | 否 | ✓ 1225ms | ✓ 1250ms | 否 | http |
| 39.104.201.40:7890 | ✓ 706ms | ✓ 1688ms | 否 | ✓ 983ms | ✓ 741ms | http |
| 116.80.49.170:3172 | ✓ 1568ms | 否 | 否 | ✓ 1975ms | ✓ 1918ms | http |
| 107.172.125.217:3128 | ✓ 594ms | 否 | ✓ 529ms | ✓ 687ms | ✓ 630ms | http |
| 61.72.110.114:3128 | ✓ 1919ms | ✓ 883ms | 否 | 否 | ✓ 1497ms | http |
| 137.184.14.135:3128 | ✓ 224ms | ✓ 744ms | ✓ 840ms | ✓ 646ms | ✓ 482ms | http |
| 61.72.221.94:3128 | ✓ 1943ms | ✓ 996ms | ✓ 908ms | ✓ 1007ms | ✓ 731ms | http |
| 45.136.198.40:3128 | ✓ 1172ms | 否 | ✓ 1815ms | 否 | ✓ 1987ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1924ms | ✓ 539ms | ✓ 1868ms | 否 | http |
| 61.72.110.54:3128 | ✓ 1948ms | 否 | ✓ 1276ms | 否 | ✓ 1940ms | http |
| 67.169.98.211:443 | ✓ 677ms | 否 | ✓ 170ms | ✓ 1564ms | 否 | http |
| 45.77.249.199:1236 | ✓ 838ms | 否 | ✓ 1223ms | ✓ 1088ms | ✓ 910ms | http |
| 121.230.9.26:1080 | 否 | ✓ 1740ms | ✓ 1019ms | ✓ 1374ms | ✓ 1293ms | http |
| 120.92.212.16:7890 | ✓ 1479ms | ✓ 991ms | ✓ 753ms | 否 | ✓ 771ms | http |
| 210.77.29.245:7890 | ✓ 788ms | ✓ 1292ms | ✓ 942ms | ✓ 1018ms | ✓ 802ms | http |
| 152.42.213.210:8080 | ✓ 1891ms | 否 | ✓ 1910ms | ✓ 1014ms | 否 | http |
| 103.183.10.203:3125 | ✓ 1372ms | 否 | 否 | ✓ 1309ms | ✓ 1347ms | http |
| 121.230.8.22:1080 | ✓ 1161ms | ✓ 1463ms | ✓ 1235ms | ✓ 1356ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1836ms | ✓ 1659ms | ✓ 1717ms | 否 | 否 | http |
| 220.170.182.39:9293 | ✓ 1105ms | ✓ 1181ms | ✓ 1087ms | ✓ 1148ms | ✓ 1151ms | http |
| 154.3.236.202:3128 | ✓ 683ms | 否 | ✓ 1280ms | ✓ 1528ms | ✓ 1138ms | http |
| 125.128.12.144:3128 | 否 | ✓ 1757ms | ✓ 1435ms | ✓ 1558ms | ✓ 1786ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1301ms | 否 | ✓ 1010ms | ✓ 1977ms | http |
| 47.101.149.27:9010 | ✓ 1080ms | ✓ 1069ms | 否 | ✓ 1229ms | 否 | http |
| 120.92.212.16:8890 | ✓ 930ms | 否 | ✓ 1458ms | 否 | ✓ 764ms | http |
| 172.105.118.164:3128 | ✓ 1916ms | 否 | ✓ 1914ms | ✓ 1580ms | ✓ 1158ms | http |
| 121.230.9.148:1080 | 否 | ✓ 1875ms | ✓ 807ms | ✓ 1254ms | ✓ 886ms | http |
| 168.235.110.63:3128 | ✓ 692ms | 否 | ✓ 1574ms | ✓ 1425ms | ✓ 1052ms | http |
| 47.105.98.23:3128 | ✓ 1150ms | ✓ 861ms | 否 | 否 | ✓ 960ms | http |
| 103.82.23.118:5196 | ✓ 1883ms | 否 | 否 | ✓ 1763ms | ✓ 1554ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1509ms | ✓ 1107ms | ✓ 1240ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1209ms | ✓ 926ms | ✓ 606ms | 否 | 否 | http |
| 194.213.18.200:443 | 否 | ✓ 1160ms | 否 | ✓ 1319ms | ✓ 1197ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1902ms | 否 | ✓ 1264ms | ✓ 1277ms | http |
| 152.42.213.210:80 | ✓ 1122ms | 否 | ✓ 990ms | ✓ 1092ms | ✓ 1981ms | http |
| 152.70.98.46:8888 | ✓ 883ms | ✓ 1365ms | ✓ 1390ms | ✓ 1087ms | ✓ 800ms | http |
| 113.132.112.110:9000 | ✓ 1209ms | ✓ 1151ms | ✓ 1078ms | ✓ 1219ms | ✓ 1404ms | http |
| 61.52.131.172:8443 | ✓ 710ms | ✓ 1023ms | ✓ 841ms | ✓ 954ms | ✓ 746ms | http |
| 121.230.8.144:1080 | ✓ 952ms | ✓ 1326ms | ✓ 939ms | ✓ 1453ms | 否 | http |
| 201.150.116.32:999 | ✓ 1174ms | 否 | ✓ 1234ms | ✓ 1587ms | 否 | http |
| 59.46.216.131:30001 | ✓ 924ms | ✓ 1044ms | ✓ 893ms | 否 | ✓ 1923ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1509ms | ✓ 1326ms | ✓ 1039ms | http |
| 43.165.195.107:3128 | ✓ 768ms | ✓ 1373ms | ✓ 1119ms | ✓ 1166ms | 否 | http |
| 106.14.203.63:3333 | ✓ 660ms | ✓ 884ms | ✓ 792ms | ✓ 814ms | ✓ 727ms | http |
| 120.55.163.237:10086 | 否 | 否 | ✓ 1725ms | ✓ 1321ms | ✓ 723ms | http |
| 103.39.51.190:8080 | ✓ 1995ms | 否 | ✓ 1159ms | ✓ 1423ms | ✓ 1272ms | http |

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
